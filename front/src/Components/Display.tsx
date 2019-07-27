import * as React from "react";

import { ReducerContext } from "../Reducer";
import { Place } from "../State";
import {
  mouseMoveAction,
  mouseClickAction,
  selectPointAction,
  mouseUpAction
} from "../Action";
import Line from "./Line";
import Point from "./Point";

type DisplayProps = {
  width: number;
  height: number;
  viewWidth: number;
  viewHeight: number;
};

export default ({ ...props }: DisplayProps) => {
  const { state, dispatcher } = React.useContext(ReducerContext);

  const placeToSvgSpace = toSvgSpace(
    [props.width, props.height],
    [props.viewWidth, props.viewHeight]
  );

  const routes = state.shortestPathKey
    ? state.shortestPaths[state.shortestPathKey][state.shortestPath]
    : null;

  return (
    <svg
      width={`${props.viewWidth}px`}
      height={`${props.viewHeight}px`}
      viewBox={`0 0 ${props.width} ${props.height}`}
      onMouseMove={e =>
        state.isClick
          ? dispatcher(
              mouseMoveAction(
                placeToSvgSpace({
                  x: e.nativeEvent.offsetX,
                  y: e.nativeEvent.offsetY
                })
              )
            )
          : null
      }
      onMouseDown={e =>
        dispatcher(
          mouseClickAction(
            placeToSvgSpace({
              x: e.nativeEvent.offsetX,
              y: e.nativeEvent.offsetY
            })
          )
        )
      }
      onMouseUp={() => dispatcher(mouseUpAction())}
      style={{ border: "1px solid black" }}
    >
      {state.roads.map(({ edge, isHighWay }, i) => (
        <Line
          key={i}
          from={state.places[edge[0]]}
          to={state.places[edge[1]]}
          width={0.05}
          isShowLength={true}
          color={isHighWay ? "#ff8844" : "black"}
        />
      ))}
      {routes
        ? routes.path.map(([from, to], i) => (
            <Line
              key={i}
              from={state.places[from]}
              to={state.places[to]}
              width={0.05}
              color="#4488ff"
              isShowLength={false}
            />
          ))
        : null}
      {state.places.map((x, i) => (
        <Point
          place={x}
          size={0.1}
          key={i}
          index={i + 1}
          onMouseDown={e => (
            e.stopPropagation(), dispatcher(selectPointAction(i, false))
          )}
          color={
            state.selectPoint &&
            !state.selectPoint.isAdded &&
            i === state.selectPoint.index
              ? "green"
              : "red"
          }
        />
      ))}
      {state.addedPlaces.map((x, i) => (
        <Point
          place={x}
          size={0.1}
          key={i}
          index={i + 1}
          onMouseDown={e => (
            e.stopPropagation(), dispatcher(selectPointAction(i, true))
          )}
          color={
            state.selectPoint &&
            state.selectPoint.isAdded &&
            i === state.selectPoint.index
              ? "green"
              : "blue"
          }
        />
      ))}
    </svg>
  );
};

const toSvgSpace = ([w, h]: [number, number], [vw, vh]: [number, number]) => ({
  x,
  y,
  ...p
}: Place): Place => ({
  ...p,
  x: (x / vw) * w,
  y: (y / vh) * h
});
