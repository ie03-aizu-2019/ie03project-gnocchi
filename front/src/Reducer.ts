import {createContext, Dispatch} from 'react';

import {Action} from './Action';
import {Place, Query, State} from './State';

const exportQuery =
    (places: Place[], roads: [number, number][], addedPlaces: Place[],
     queries: Query[]): string => {
      const placesStr = places.map(p => `${p.x} ${p.y}`).join('\n');
      const roadsStr = roads.map(([f, t]) => `${f + 1} ${t + 1}`).join('\n');
      const addedPlacesStr = addedPlaces.map(p => `${p.x} ${p.y}`).join('\n');
      const queriesStr =
          queries.map(({start, end, num}) => `${start} ${end} ${num}`)
              .join('\n');

      return `${places.length} ${roads.length} ${addedPlaces.length} ${
          queries.length}
${[placesStr, roadsStr, addedPlacesStr, queriesStr].join('\n')}`;
    };

const importQuery = (query: string): [Place[], [number, number][], Query[]] => {
  const lines = query.split('\n');
  const [p, r, a, q] = lines[0].split(' ').map(Number);

  const places = lines.slice(1, p + 1)
                     .map(x => x.split(' ').map(Number))
                     .map(([x, y]) => ({x, y, isAdded: false}));

  const roads = lines.slice(p + 1, r + p + 1)
                    .map(x => x.split(' ').map(Number))
                    .map(([f, t]) => [f - 1, t - 1] as [number, number]);

  const addedPlaces = lines.slice(r + p + 1, r + p + a + 1)
                          .map(x => x.split(' ').map(Number))
                          .map(([x, y]) => ({x, y, isAdded: true}));

  const queries = lines.slice(r + p + a + 1, r + p + a + q + 1)
                      .map(x => x.split(' '))
                      .map(([s, e, n]) => ({start: s, end: e, num: Number(n)}));

  return [places.concat(addedPlaces), roads, queries];
};

export default (state: State, action: Action): State => {
  switch (action.type_) {
    case 'ChangeMode':
      return {...state, mode: action.mode, shortestPathKey: null};

    case 'InputQuery':
      return {...state, testQuery: action.query};

    case 'Export':
      return {
        ...state,
        testQuery: exportQuery(
            state.places.filter(x => !x.isAdded), state.roads,
            state.places.filter(x => x.isAdded), state.queries)
      };

    case 'Import': {
      const [places, roads] = importQuery(state.testQuery);
      return {...state, places, roads};
    }

    case 'ToAPI': {
      const [places, roads] = importQuery(action.newQuery);
      return {...state, places, roads, testQuery: action.newQuery};
    }

    case 'SelectShortestPaths': {
      return {...state, shortestPathKey: action.key};
    }

    case 'SelectShortestPath': {
      return {...state, shortestPath: action.index};
    }

    // Modeで挙動を変える
    default:
      return modableReduser(state, action);
  }
};

const modableReduser = (state: State, action: Action): State => {
  if (state.mode === 'Add') {
    switch (action.type_) {
      case 'MouseClick': {
        const places =
            [...state.places, {x: action.x, y: action.y, isAdded: false}];
        const testQuery = exportQuery(
            places.filter(x => !x.isAdded), state.roads,
            places.filter(x => x.isAdded), state.queries);

        return {...state, places, testQuery};
      }

      case 'SelectPoint': {
        const places = state.places.map(
            (x, i) =>
                i === action.pointIndex ? {...x, isAdded: !x.isAdded} : x);
        const testQuery = exportQuery(
            places.filter(x => !x.isAdded), state.roads,
            places.filter(x => x.isAdded), state.queries);

        return {...state, places, testQuery};
      }
    }
  } else if (state.mode === 'Remove') {
    switch (action.type_) {
      case 'SelectPoint': {
        const places = state.places.filter((_, i) => i !== action.pointIndex);
        const roads = state.roads
                          .filter(
                              x => !x.map(y => y === action.pointIndex)
                                        .reduce((acc, y) => acc || y))
                          .map(
                              x => x.map(
                                       y =>
                                           (y > action.pointIndex ?
                                                y - 1 :
                                                y)) as [number, number]);
        const testQuery = exportQuery(
            places.filter(x => !x.isAdded), roads,
            places.filter(x => x.isAdded), state.queries);

        return {...state, places, roads, testQuery};
      }
    }
  } else if (state.mode === 'DrawLine') {
    switch (action.type_) {
      case 'SelectPoint':
        if (state.selectPoint !== null) {
          const roads = [
            ...state.roads,
            [state.selectPoint, action.pointIndex] as [number, number]
          ];
          const testQuery = exportQuery(
              state.places.filter(x => !x.isAdded), roads,
              state.places.filter(x => x.isAdded), state.queries);

          return {...state, selectPoint: null, roads, testQuery};
        } else {
          return {...state, selectPoint: action.pointIndex};
        }

      case 'MouseClick':
        return {...state, selectPoint: null};
    }
  } else if (state.mode === 'Move') {
    switch (action.type_) {
      case 'SelectPoint':
        return {...state, isClick: true, selectPoint: action.pointIndex};
      case 'MouseUp':
        return {...state, isClick: false, selectPoint: null};
      case 'MouseMove':
        const places = state.places.map(
            (x, i) =>
                state.selectPoint === i ? {...x, x: action.x, y: action.y} : x);
        const testQuery = exportQuery(
            places.filter(x => !x.isAdded), state.roads,
            places.filter(x => x.isAdded), state.queries);

        return {...state, places, testQuery};
    }
  }

  return state;
};

type Reducer = {
  state: State; dispatcher: Dispatch<Action>;
};
export const ReducerContext = createContext<Reducer>(null);
