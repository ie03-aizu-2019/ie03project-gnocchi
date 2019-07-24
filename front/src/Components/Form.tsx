import * as React from "react";

import { Mode, Action, changeModeAction } from "../Reducer";

import Button from "./Button";
import Grid from "./Grid";

type FormProps = {
  mode: Mode;
  dispatcher: (arg0: Action) => void;
};

export default ({ mode, dispatcher }: FormProps) => (
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
