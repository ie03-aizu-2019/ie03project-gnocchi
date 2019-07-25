import * as React from "react";
import styled from "styled-components";

import {
  importAction,
  exportAction,
  inputQueryAction,
  ReducerContext,
  toAPIAction,
  Action
} from "../Reducer";

import Button from "./Button";
import Grid from "./Grid";
import Caller, { EnumCrossPoints, RecomendCrossPoints } from "../ApiCall";

type IOProps = {
  query: string;
};

const enumCrossPointsCall = async (
  query: string,
  dispatcher: React.Dispatch<Action>
) => {
  const newQuery = await Caller(EnumCrossPoints, {
    method: "POST",
    headers: { "Content-Type": "text/plain" },
    body: query
  });

  dispatcher(toAPIAction(newQuery));
};

const recomendCrossPointsCall = async (
  query: string,
  dispatcher: React.Dispatch<Action>
) => {
  const newQuery = await Caller(RecomendCrossPoints, {
    method: "POST",
    headers: { "Content-Type": "text/plain" },
    body: query
  });

  dispatcher(toAPIAction(newQuery));
};

export default ({ query }: IOProps) => {
  const { dispatcher } = React.useContext(ReducerContext);

  return (
    <Grid rows={["32px", "32px", "32px", "1fr"]} columns={["1fr"]} gap="8px">
      <Button onClick={() => enumCrossPointsCall(query, dispatcher)}>
        CrossPoints
      </Button>
      <Button onClick={() => recomendCrossPointsCall(query, dispatcher)}>
        RecomendCrossPoints
      </Button>
      <Button onClick={() => dispatcher(importAction())}>Inport</Button>
      <Textarea
        value={query}
        onChange={e => dispatcher(inputQueryAction(e.currentTarget.value))}
        style={{ fontSize: "16px" }}
      />
    </Grid>
  );
};

const Textarea = styled.textarea`
  box-sizing: border-box;
`;
