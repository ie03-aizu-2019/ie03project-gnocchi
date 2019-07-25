import * as React from "react";

import {
  Place,
  Action,
  mouseMoveAction,
  mouseClickAction,
  selectPointAction,
  mouseUpAction,
  ReducerContext
} from "../Reducer";
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
                  y: e.nativeEvent.offsetY,
                  isAdded: false
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
              y: e.nativeEvent.offsetY,
              isAdded: false
            })
          )
        )
      }
      onMouseUp={() => dispatcher(mouseUpAction())}
      style={{ border: "1px solid black" }}
    >
      {state.roads.map(([from, to], i) => (
        <Line
          key={i}
          from={state.places[from]}
          to={state.places[to]}
          width={0.05}
        />
      ))}
      {state.places.map((x, i) => (
        <Point
          place={x}
          size={0.1}
          key={i}
          index={i + 1}
          onMouseDown={e => (
            e.stopPropagation(), dispatcher(selectPointAction(i))
          )}
          color={
            x.isAdded ? "blue" : state.selectPoint === i ? "green" : undefined
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
