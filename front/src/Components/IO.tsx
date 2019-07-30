import * as React from "react";
import styled from "styled-components";

import { ReducerContext } from "../Reducer";
import {
  importAction,
  inputQueryAction,
  toAPIAction,
  randomAction,
  Action,
  detectionHighWaysAction,
  shortestPathsAction
} from "../Action";

import Button from "./Button";
import Grid from "./Grid";
import Caller, {
  EnumCrossPoints,
  RecomendCrossPoints,
  DetectionHighWays,
  ShortestPath
} from "../ApiCall";
import { Route } from "../State";

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

  dispatcher(toAPIAction(newQuery.trim()));
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

  dispatcher(toAPIAction(newQuery.trim()));
};

const detectionHighWaysCall = async (
  query: string,
  dispatcher: React.Dispatch<Action>
) => {
  const highWays = await Caller(DetectionHighWays, {
    method: "POST",
    headers: { "Content-Type": "text/plain" },
    body: query
  }).then((x: string) => JSON.parse(x) as string[][]);

  dispatcher(
    detectionHighWaysAction(
      highWays
        .map(xs => xs.map(Number))
        .map(([f, t]) => [f, t] as [number, number])
    )
  );
};

const shortestPathsCall = async (
  query: string,
  dispatcher: React.Dispatch<Action>
) => {
  const shortestPaths = await Caller(ShortestPath, {
    method: "POST",
    headers: { "Content-Type": "text/plain" },
    body: query
  }).then(
    (x: string) =>
      JSON.parse(x) as {
        paths: { [key: string]: [string, string][][] };
        query: string;
      }
  );

  dispatcher(toAPIAction(shortestPaths.query.trim()));
  dispatcher(
    shortestPathsAction(
      Object.keys(shortestPaths.paths).reduce((acc, x) => {
        return {
          ...acc,
          [x]: shortestPaths.paths[x].map(
            ys =>
              ({
                path: ys.map(
                  ([f, t]) => [Number(f) - 1, Number(t) - 1] as [number, number]
                )
              } as Route)
          )
        };
      }, {})
    )
  );
};

export default ({ query }: IOProps) => {
  const { dispatcher } = React.useContext(ReducerContext);

  return (
    <Grid
      rows={["32px", "32px", "32px", "32px", "32px", "32px", "1fr"]}
      columns={["1fr"]}
      gap="8px"
    >
      <Button onClick={() => enumCrossPointsCall(query, dispatcher)}>
        CrossPoints
      </Button>
      <Button onClick={() => recomendCrossPointsCall(query, dispatcher)}>
        RecomendCrossPoints
      </Button>
      <Button onClick={() => detectionHighWaysCall(query, dispatcher)}>
        Highway detection
      </Button>
      <Button onClick={() => shortestPathsCall(query, dispatcher)}>
        ShortestPaths
      </Button>
      <Button onClick={() => dispatcher(randomAction())}>Random</Button>
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
