import * as React from "react";

import { ReducerContext } from "../Reducer";
import { Mode } from "../State";
import {
  changeModeAction,
  selectShortestPath,
  selectShortestPaths
} from "../Action";
import Button from "./Button";
import Grid from "./Grid";

type FormProps = {
  mode: Mode;
};

export default ({ mode }: FormProps) => {
  const { state, dispatcher } = React.useContext(ReducerContext);

  return (
    <Grid rows={["1fr", "1fr"]} columns={["1fr"]}>
      <Grid rows={["repeat(6, 32px)"]} columns={["1fr"]} gap="8px">
        {([
          "Add",
          "AddedPoint",
          "Remove",
          "DrawLine",
          "Move",
          "ShowPath"
        ] as Mode[]).map((x, i) => {
          return (
            <Button
              key={i}
              onClick={() => dispatcher(changeModeAction(x))}
              disabled={mode === x}
            >
              {x}
            </Button>
          );
        })}
      </Grid>
      {mode === "ShowPath" ? (
        <Grid
          rows={Object.keys(state.shortestPaths)
            .map(() => "32px")
            .concat(["32px"])}
          columns={["1fr"]}
          gap="8px"
        >
          <Grid rows={["1fr"]} columns={["32px", "1fr"]}>
            <span>{state.shortestPath}</span>
            <input
              type="range"
              min={0}
              max={
                state.shortestPathKey
                  ? state.shortestPaths[state.shortestPathKey].length - 1
                  : 0
              }
              onInput={e =>
                dispatcher(selectShortestPath(Number(e.currentTarget.value)))
              }
            />
          </Grid>
          {Object.keys(state.shortestPaths).map((x, i) => (
            <Button
              key={i}
              disabled={x === state.shortestPathKey}
              onClick={() => dispatcher(selectShortestPaths(x))}
            >
              {x}
            </Button>
          ))}
        </Grid>
      ) : (
        ""
      )}
    </Grid>
  );
};
