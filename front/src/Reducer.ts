import { createContext, Dispatch } from "react";

import { Action } from "./Action";
import { Place, Road, Query, State } from "./State";

const exportQuery = (
  places: Place[],
  roads: Road[],
  addedPlaces: Place[],
  queries: Query[]
): string => {
  const placesStr = places.map(p => `${p.x} ${p.y}`).join("\n");
  const roadsStr = roads
    .map(x => x.edge)
    .map(([f, t]) => `${f + 1} ${t + 1}`)
    .join("\n");
  const addedPlacesStr = addedPlaces.map(p => `${p.x} ${p.y}`).join("\n");
  const queriesStr = queries
    .map(({ start, end, num }) => `${start} ${end} ${num}`)
    .join("\n");

  return `${places.length} ${roads.length} ${addedPlaces.length} ${
    queries.length
  }
${[placesStr, roadsStr, addedPlacesStr, queriesStr].join("\n")}`;
};

const importQuery = (query: string): [Place[], Road[], Place[], Query[]] => {
  const lines = query.split("\n");
  const [p, r, a, q] = lines[0].split(" ").map(Number);

  const places = lines
    .slice(1, p + 1)
    .map(x => x.split(" ").map(Number))
    .map(([x, y]) => ({ x, y }));

  const roads = lines
    .slice(p + 1, r + p + 1)
    .map(x => x.split(" ").map(Number))
    .map(([f, t]) => [f - 1, t - 1] as [number, number])
    .map(x => ({ edge: x, isHighWay: false }));

  const addedPlaces = lines
    .slice(r + p + 1, r + p + a + 1)
    .map(x => x.split(" ").map(Number))
    .map(([x, y]) => ({ x, y }));

  const queries = lines
    .slice(r + p + a + 1, r + p + a + q + 1)
    .map(x => x.split(" "))
    .map(([s, e, n]) => ({ start: s, end: e, num: Number(n) }));

  return [places, roads, addedPlaces, queries];
};

export default (state: State, action: Action): State => {
  switch (action.type_) {
    case "ChangeMode":
      return { ...state, mode: action.mode, shortestPathKey: null };

    case "InputQuery":
      return { ...state, testQuery: action.query };

    case "Export":
      return {
        ...state,
        testQuery: exportQuery(
          state.places,
          state.roads,
          state.addedPlaces,
          state.queries
        )
      };

    case "Import": {
      const [places, roads, addedPlaces, queries] = importQuery(
        state.testQuery
      );
      return { ...state, places, roads, addedPlaces, queries };
    }

    case "ToAPI": {
      const [places, roads, addedPlaces, queries] = importQuery(
        action.newQuery
      );
      return {
        ...state,
        places,
        roads,
        addedPlaces,
        queries,
        testQuery: action.newQuery
      };
    }

    case "SelectShortestPaths": {
      return { ...state, shortestPathKey: action.key };
    }

    case "SelectShortestPath": {
      return { ...state, shortestPath: action.index };
    }

    case "Random": {
      const N = 10;

      const places = [...Array(N).fill(null)].map(
        () =>
          ({
            x: Math.random() * 10,
            y: Math.random() * 10
          } as Place)
      );

      const addedPlaces = [...Array(Math.round(N / 2)).fill(null)].map(
        () =>
          ({
            x: Math.random() * 10,
            y: Math.random() * 10
          } as Place)
      );

      const roads = [...Array(N).fill(null)]
        .map((_, i) => {
          const r = [] as Road[];
          for (let n = 0; n < N - i - 1; n++) {
            if (Math.random() < 0.5) {
              r.push({ edge: [i, n + i + 1], isHighWay: false });
            }
          }
          return r;
        })
        .reduce((acc, x) => acc.concat(x), []);

      const testQuery = exportQuery(places, roads, addedPlaces, state.queries);

      return {
        ...state,
        places,
        roads,
        addedPlaces,
        testQuery
      };
    }

    case "DetectionHighWays": {
      const roads = state.roads.map(({ edge }) => ({
        edge,
        isHighWay: action.highWays
          .map(([f, t]) => [f - 1, t - 1])
          .some(
            ([f, t]) =>
              (edge[0] === f && edge[1] === t) ||
              (edge[1] === f && edge[0] === t)
          )
      }));

      return { ...state, roads };
    }

    case "ShortestPaths":
      return { ...state, shortestPaths: action.shortestPaths };

    // Modeで挙動を変える
    default:
      return modableReduser(state, action);
  }
};

const modableReduser = (state: State, action: Action): State => {
  if (state.mode === "Add") {
    switch (action.type_) {
      case "MouseClick": {
        const places = [...state.places, { x: action.x, y: action.y }];
        const testQuery = exportQuery(
          places,
          state.roads,
          state.addedPlaces,
          state.queries
        );

        return { ...state, places, testQuery };
      }
    }
  } else if (state.mode === "AddedPoint") {
    switch (action.type_) {
      case "MouseClick": {
        const addedPlaces = [
          ...state.addedPlaces,
          { x: action.x, y: action.y }
        ];
        const testQuery = exportQuery(
          state.places,
          state.roads,
          addedPlaces,
          state.queries
        );

        return { ...state, addedPlaces, testQuery };
      }
    }
  } else if (state.mode === "Remove") {
    switch (action.type_) {
      case "SelectPoint": {
        const addedPlaces = action.isAdded
          ? state.addedPlaces.filter((_, i) => i !== action.pointIndex)
          : state.addedPlaces;
        const places = !action.isAdded
          ? state.places.filter((_, i) => i !== action.pointIndex)
          : state.places;

        const roads = !action.isAdded
          ? state.roads
              .filter(
                x =>
                  !x.edge
                    .map(y => y === action.pointIndex)
                    .reduce((acc, y) => acc || y)
              )
              .map(x => ({
                ...x,
                edge: x.edge.map(y => (y > action.pointIndex ? y - 1 : y)) as [
                  number,
                  number
                ]
              }))
          : state.roads;

        const testQuery = exportQuery(
          places,
          roads,
          addedPlaces,
          state.queries
        );

        return { ...state, places, roads, addedPlaces, testQuery };
      }
    }
  } else if (state.mode === "DrawLine") {
    switch (action.type_) {
      case "SelectPoint":
        if (action.isAdded) return state;
        if (
          state.selectPoint !== null &&
          state.selectPoint.index !== action.pointIndex
        ) {
          const roads = [
            ...state.roads,
            {
              edge: [state.selectPoint.index, action.pointIndex] as [
                number,
                number
              ],
              isHighWay: false
            }
          ];
          const testQuery = exportQuery(
            state.places,
            roads,
            state.addedPlaces,
            state.queries
          );

          return { ...state, selectPoint: null, roads, testQuery };
        } else {
          return {
            ...state,
            selectPoint: { index: action.pointIndex, isAdded: action.isAdded }
          };
        }

      case "MouseClick":
        return { ...state, selectPoint: null };
    }
  } else if (state.mode === "Move") {
    switch (action.type_) {
      case "SelectPoint":
        return {
          ...state,
          isClick: true,
          selectPoint: { index: action.pointIndex, isAdded: action.isAdded }
        };
      case "MouseUp":
        return { ...state, isClick: false, selectPoint: null };
      case "MouseMove":
        const places = !state.selectPoint.isAdded
          ? state.places.map((x, i) =>
              state.selectPoint.index === i
                ? { ...x, x: action.x, y: action.y }
                : x
            )
          : state.places;

        const addedPlaces = state.selectPoint.isAdded
          ? state.addedPlaces.map((x, i) =>
              state.selectPoint.index === i
                ? { ...x, x: action.x, y: action.y }
                : x
            )
          : state.addedPlaces;

        const testQuery = exportQuery(
          places,
          state.roads,
          addedPlaces,
          state.queries
        );

        return { ...state, places, addedPlaces, testQuery };
    }
  }

  return state;
};

type Reducer = {
  state: State;
  dispatcher: Dispatch<Action>;
};
export const ReducerContext = createContext<Reducer>(null);
