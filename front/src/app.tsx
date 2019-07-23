import * as React from "react";
import { render } from "react-dom";

import reduser, { init } from "./Reducer";
import Form from "./Components/Form";
import Display from "./Components/Display";

const mount = document.getElementById("mount");

type AppProps = {
  width: number;
  height: number;
};

const App = ({ width, height }: AppProps) => {
  const [state, dispatcher] = React.useReducer(reduser, init);

  return (
    <>
      <Display
        width={width}
        height={height}
        {...state}
        dispatcher={dispatcher}
      />
      <Form {...state} dispatcher={dispatcher} query={state.testQuery} />
    </>
  );
};

render(<App width={10} height={10} />, mount);
