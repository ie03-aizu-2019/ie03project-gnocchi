export type Mode =
  | "Add"
  | "AddedPoint"
  | "DrawLine"
  | "Remove"
  | "Move"
  | "ShowPath";

export type Place = {
  x: number;
  y: number;
};

export type Road = {
  edge: [number, number];
  isHighWay: boolean;
};

export type Query = {
  start: string;
  end: string;
  num: number;
};

export type Route = {
  path: [number, number][];
};

export type State = {
  mode: Mode;
  selectPoint: { index: number; isAdded: boolean } | null;
  isClick: boolean;
  shortestPathKey: string | null;
  shortestPath: number;
  testQuery: string;
  places: Place[];
  roads: Road[];
  addedPlaces: Place[];
  queries: Query[];
  shortestPaths: { [key: string]: Route[] };
};

export const init: State = {
  mode: "Add",
  testQuery: "",
  selectPoint: null,
  isClick: false,
  shortestPathKey: null,
  shortestPath: 0,
  places: [],
  roads: [],
  addedPlaces: [],
  queries: [],
  shortestPaths: {
    "1 2 1": [{ path: [[0, 1]] }, { path: [[0, 2]] }],
    "1 4 1": [{ path: [[0, 2]] }]
  }
};
