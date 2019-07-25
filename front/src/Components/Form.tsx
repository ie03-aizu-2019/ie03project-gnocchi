import * as React from "react";

import { Mode, changeModeAction, ReducerContext } from "../Reducer";
import Button from "./Button";
import Grid from "./Grid";

type FormProps = {
  mode: Mode;
};

export default ({ mode }: FormProps) => {
  const { dispatcher } = React.useContext(ReducerContext);

  return (
    <Grid rows={["repeat(4, 32px)"]} columns={["1f"]} gap="8px">
      {(["Add", "Remove", "DrawLine", "Move"] as Mode[]).map((x, i) => {
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
  );
};
