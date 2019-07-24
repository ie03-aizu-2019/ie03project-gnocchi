import * as React from "react";
import { render } from "react-dom";
import styled from "styled-components";

import reduser, { init } from "./Reducer";
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
    <Grid
      rows={["1fr", "1fr"]}
      columns={["1fr", `${props.viewWidth}px`, "1fr"]}
      gap="16px"
    >
      <Form {...state} dispatcher={dispatcher} />
      <Display {...props} {...state} dispatcher={dispatcher} />
      <IO query={state.testQuery} dispatcher={dispatcher} />
    </Grid>
  );
};

render(<App width={10} height={10} viewWidth={640} viewHeight={640} />, mount);
