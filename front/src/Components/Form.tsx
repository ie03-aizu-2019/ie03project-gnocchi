import * as React from "react";

import {
  Mode,
  Action,
  changeModeAction,
  inputQueryAction,
  exportAction,
  importAction
} from "../Reducer";

type FormProps = {
  mode: Mode;
  query: string;
  dispatcher: (arg0: Action) => void;
};

export default ({ mode, query, dispatcher }: FormProps) => {
  return (
    <div>
      {(["Add", "Remove", "DrawLine", "Move"] as Mode[]).map((x, i) => {
        return (
          <button
            key={i}
            onClick={() => dispatcher(changeModeAction(x))}
            disabled={mode === x}
          >
            {x}
          </button>
        );
      })}
      <div>
        <button onClick={() => console.log("API と通信")}>SendData</button>
      </div>
      <div>
        <button onClick={() => dispatcher(importAction())}>Inport</button>
        <button onClick={() => dispatcher(exportAction())}>Export</button>
      </div>
      <div>
        <textarea
          value={query}
          onChange={e => dispatcher(inputQueryAction(e.currentTarget.value))}
        />
      </div>
    </div>
  );
};
