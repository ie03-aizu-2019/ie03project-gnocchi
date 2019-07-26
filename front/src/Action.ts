import {Mode, Place} from './State';

type MouseMove = {
  readonly type_: 'MouseMove'; x: number; y: number;
};

type MouseClick = {
  readonly type_: 'MouseClick'; x: number; y: number;
};

type MouseUp = {
  readonly type_: 'MouseUp';
};

type SelectPoint = {
  readonly type_: 'SelectPoint'; pointIndex: number;
};

type ChangeMode = {
  readonly type_: 'ChangeMode'; mode: Mode;
};

type InputQuery = {
  readonly type_: 'InputQuery'; query: string;
};

type Export = {
  readonly type_: 'Export';
};

type Import = {
  readonly type_: 'Import';
};

type ToAPI = {
  readonly type_: 'ToAPI'; newQuery: string;
};

type SelectShortestPaths = {
  readonly type_: 'SelectShortestPaths'; key: string;
};

type SelectShortestPath = {
  readonly type_: 'SelectShortestPath'; index: number;
}

export type Action =|MouseMove|MouseUp|MouseClick|SelectPoint|ChangeMode|
    InputQuery|Export|Import|ToAPI|SelectShortestPaths|SelectShortestPath;

export const mouseMoveAction = (p: Place): MouseMove =>
    ({type_: 'MouseMove', ...p});

export const mouseClickAction = (p: Place): MouseClick =>
    ({type_: 'MouseClick', ...p});

export const mouseUpAction = (): MouseUp => ({type_: 'MouseUp'});

export const selectPointAction = (n: number): SelectPoint =>
    ({type_: 'SelectPoint', pointIndex: n});

export const changeModeAction = (mode: Mode): ChangeMode =>
    ({type_: 'ChangeMode', mode: mode});

export const inputQueryAction = (q: string): InputQuery =>
    ({type_: 'InputQuery', query: q});

export const exportAction = (): Export => ({type_: 'Export'});

export const importAction = (): Import => ({type_: 'Import'});

export const toAPIAction = (q: string): ToAPI =>
    ({type_: 'ToAPI', newQuery: q});

export const selectShortestPaths = (key: string): SelectShortestPaths =>
    ({type_: 'SelectShortestPaths', key});

export const selectShortestPath = (index: number): SelectShortestPath =>
    ({type_: 'SelectShortestPath', index});
