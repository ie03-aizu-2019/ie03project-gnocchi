export type Mode = "Add" | "DrawLine" | "Remove" | "Move";

export type Place = {
  x: number;
  y: number;
};

export type State = {
  mode: Mode;

  selectPoint: number | null;
  isClick: boolean;

  testQuery: string;

  places: Place[];
  roads: [number, number][];
};

type MouseMove = {
  readonly type_: "MouseMove";
  x: number;
  y: number;
};

type MouseClick = {
  readonly type_: "MouseClick";
  x: number;
  y: number;
};

type MouseUp = {
  readonly type_: "MouseUp";
};

type SelectPoint = {
  readonly type_: "SelectPoint";
  pointIndex: number;
};

type ChangeMode = {
  readonly type_: "ChangeMode";
  mode: Mode;
};

type InputQuery = {
  readonly type_: "InputQuery";
  query: string;
};

type Export = {
  readonly type_: "Export";
};

type Import = {
  readonly type_: "Import";
};

export type Action =
  | MouseMove
  | MouseUp
  | MouseClick
  | SelectPoint
  | ChangeMode
  | InputQuery
  | Export
  | Import;

export const init: State = {
  mode: "Add",
  testQuery: "",
  selectPoint: null,
  isClick: false,
  places: [],
  roads: []
};

export const mouseMoveAction = (p: Place): MouseMove => ({
  type_: "MouseMove",
  ...p
});

export const mouseClickAction = (p: Place): MouseClick => ({
  type_: "MouseClick",
  ...p
});

export const mouseUpAction = (): MouseUp => ({
  type_: "MouseUp"
});

export const selectPointAction = (n: number): SelectPoint => ({
  type_: "SelectPoint",
  pointIndex: n
});

export const changeModeAction = (mode: Mode): ChangeMode => ({
  type_: "ChangeMode",
  mode: mode
});

export const inputQueryAction = (q: string): InputQuery => ({
  type_: "InputQuery",
  query: q
});

export const exportAction = (): Export => ({ type_: "Export" });

export const importAction = (): Import => ({ type_: "Import" });

const exportQuery = (places: Place[], roads: [number, number][]): string => {
  const placesStr = places.map(p => `${p.x} ${p.y}`).join("\n");
  const roadsStr = roads.map(([f, t]) => `${f + 1} ${t + 1}`).join("\n");

  return `${places.length} ${roads.length} 0 0
${placesStr}
${roadsStr}`;
};

const importQuery = (query: string): [Place[], [number, number][]] => {
  const lines = query.split("\n");
  const [p, r] = lines[0].split(" ").map(Number);
  const places = lines
    .slice(1, p + 1)
    .map(x => x.split(" ").map(Number))
    .map(([x, y]) => ({ x, y }));

  const roads = lines
    .slice(p + 1, r + p + 1)
    .map(x => x.split(" ").map(Number))
    .map(([f, t]) => [f - 1, t - 1] as [number, number]);

  return [places, roads];
};

export default (state: State, action: Action): State => {
  switch (action.type_) {
    case "ChangeMode":
      return { ...state, mode: action.mode };

    case "InputQuery":
      return { ...state, testQuery: action.query };

    case "Export":
      return { ...state, testQuery: exportQuery(state.places, state.roads) };

    case "Import":
      const [places, roads] = importQuery(state.testQuery);
      return { ...state, places, roads };

    // Modeで挙動を変える
    default:
      return modableReduser(state, action);
  }
};

const modableReduser = (state: State, action: Action): State => {
  if (state.mode === "Add") {
    switch (action.type_) {
      case "MouseClick":
        return {
          ...state,
          places: [...state.places, { x: action.x, y: action.y }]
        };
    }
  } else if (state.mode === "Remove") {
    switch (action.type_) {
      case "SelectPoint":
        return {
          ...state,
          places: state.places.filter((_, i) => i !== action.pointIndex),
          roads: state.roads
            .filter(
              x =>
                !x
                  .map(y => y === action.pointIndex)
                  .reduce((acc, y) => acc || y)
            )
            .map(
              x =>
                x.map(y => (y > action.pointIndex ? y - 1 : y)) as [
                  number,
                  number
                ]
            )
        };
    }
  } else if (state.mode === "DrawLine") {
    switch (action.type_) {
      case "SelectPoint":
        if (state.selectPoint !== null) {
          return {
            ...state,
            selectPoint: null,
            roads: [...state.roads, [state.selectPoint, action.pointIndex]]
          };
        } else {
          return {
            ...state,
            selectPoint: action.pointIndex
          };
        }

      case "MouseClick":
        return {
          ...state,
          selectPoint: null
        };
    }
  } else if (state.mode === "Move") {
    switch (action.type_) {
      case "SelectPoint":
        return {
          ...state,
          isClick: true,
          selectPoint: action.pointIndex
        };
      case "MouseUp":
        return {
          ...state,
          isClick: false,
          selectPoint: null
        };
      case "MouseMove":
        return {
          ...state,
          places: state.places.map((x, i) =>
            state.selectPoint === i ? { x: action.x, y: action.y } : x
          )
        };
    }
  }

  return state;
};
