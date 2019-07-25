import * as React from "react";
import { render } from "react-dom";

import reduser, { init, ReducerContext } from "./Reducer";
import Form from "./Components/Form";
import Display from "./Components/Display";
import IO from "./Components/IO";
import Grid from "./Components/Grid";

const mount = document.getElementById("mount");

type AppProps = {
  width: number;
  height: number;
  viewWidth: number;
  viewHeight: number;
};

const App = (props: AppProps) => {
  const [state, dispatcher] = React.useReducer(reduser, init);

  return (
    <ReducerContext.Provider value={{ state, dispatcher }}>
      <Grid
        rows={["1fr"]}
        columns={["1fr", `${props.viewWidth}px`, "1fr"]}
        gap="16px"
      >
        <Form mode={state.mode} />
        <Display {...props} />
        <IO query={state.testQuery} />
      </Grid>
    </ReducerContext.Provider>
  );
};

render(<App width={10} height={10} viewWidth={640} viewHeight={640} />, mount);
