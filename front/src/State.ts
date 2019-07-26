export type Mode = 'Add'|'DrawLine'|'Remove'|'Move'|'ShowPath';

export type Place = {
  x: number; y: number; isAdded: boolean;
};

export type Query = {
  start: string; end: string; num: number;
};

export type Route = {
  path: [number, number][];
};

export type State = {
  mode: Mode;

  selectPoint: number | null;
  isClick: boolean;
  shortestPathKey: string | null;
  shortestPath: number;

  testQuery: string;

  places: Place[];
  roads: [number, number][];
  queries: Query[];
  shortestPaths: {[key: string]: Route[]};
};

export const init: State = {
  mode: 'Add',
  testQuery: '',
  selectPoint: null,
  isClick: false,
  shortestPathKey: null,
  shortestPath: 0,
  places: [],
  roads: [],
  queries: [],
  shortestPaths: {
    '1 2 1': [{path: [[0, 1]]}, {path: [[0, 2]]}],
    '1 4 1': [{path: [[0, 2]]}]
  },
};
