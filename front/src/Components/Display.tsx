import * as React from "react";

import {
  Place,
  Action,
  mouseMoveAction,
  mouseClickAction,
  selectPointAction,
  mouseUpAction
} from "../Reducer";
import Line from "./Line";
import Point from "./Point";

type DisplayProps = {
  width: number;
  height: number;
  viewWidth: number;
  viewHeight: number;
  places: Place[];
  roads: [number, number][];
  selectPos?: Place;
  selectPoint?: number;
  isClick: boolean;
  dispatcher: (arg0: Action) => void;
};

export default ({ dispatcher, ...props }: DisplayProps) => {
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
        props.isClick
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
      {props.roads.map(([from, to], i) => (
        <Line
          key={i}
          from={props.places[from]}
          to={props.places[to]}
          width={0.05}
        />
      ))}
      {props.places.map((x, i) => (
        <Point
          place={x}
          size={0.1}
          key={i}
          index={i}
          onMouseDown={e => (
            e.stopPropagation(), dispatcher(selectPointAction(i))
          )}
          color={props.selectPoint === i ? "green" : undefined}
        />
      ))}
      {props.selectPos ? (
        <Point place={props.selectPos} size={0.1} color="blue" />
      ) : (
        ""
      )}
    </svg>
  );
};

const toSvgSpace = ([w, h]: [number, number], [vw, vh]: [number, number]) => ({
  x,
  y
}: Place): Place => ({
  x: (x / vw) * w,
  y: (y / vh) * h
});
