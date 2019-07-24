import * as React from "react";
import styled from "styled-components";

import {
  Action,
  importAction,
  exportAction,
  inputQueryAction
} from "../Reducer";

import Button from "./Button";
import FlexBox from "./FlexBox";
import Grid from "./Grid";

type IOProps = {
  query: string;
  dispatcher: (arg0: Action) => void;
};

export default ({ query, dispatcher }: IOProps) => {
  return (
    <Grid rows={["32px", "32px", "1fr"]} columns={["1fr", "1fr"]} gap="8px">
      <Button
        onClick={() => console.log("API と通信")}
        style={{ gridRow: "1", gridColumn: "1 / 3" }}
      >
        SendData
      </Button>
      <Button
        onClick={() => dispatcher(importAction())}
        style={{ gridRow: "2", gridColumn: "1" }}
      >
        Inport
      </Button>
      <Button
        onClick={() => dispatcher(exportAction())}
        style={{ gridRow: "2", gridColumn: "2" }}
      >
        Export
      </Button>
      <Textarea
        value={query}
        onChange={e => dispatcher(inputQueryAction(e.currentTarget.value))}
        style={{ gridRow: "3", gridColumn: "1/3", fontSize: "16px" }}
      />
    </Grid>
  );
};

const Textarea = styled.textarea`
  box-sizing: border-box;
`;
