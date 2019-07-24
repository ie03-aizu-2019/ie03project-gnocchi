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
  places: Place[];
  roads: [number, number][];
  selectPos?: Place;
  selectPoint?: number;
  isClick: boolean;
  dispatcher: (arg0: Action) => void;
};

export default ({ dispatcher, ...props }: DisplayProps) => {
  return (
    <svg
      width="300px"
      height="300px"
      viewBox={`0 0 ${props.width} ${props.height}`}
      onMouseMove={e =>
        props.isClick
          ? dispatcher(
              mouseMoveAction({
                x: (props.width * e.nativeEvent.offsetX) / 300,
                y: (props.height * e.nativeEvent.offsetY) / 300
              })
            )
          : null
      }
      onMouseDown={e =>
        dispatcher(
          mouseClickAction({
            x: (props.width * e.nativeEvent.offsetX) / 300,
            y: (props.height * e.nativeEvent.offsetY) / 300
          })
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
          {...x}
          size={0.1}
          key={i}
          onMouseDown={e => (
            e.stopPropagation(), dispatcher(selectPointAction(i))
          )}
          color={props.selectPoint === i ? "green" : undefined}
        />
      ))}
      {props.selectPos ? (
        <Point {...props.selectPos} size={0.1} color="blue" />
      ) : (
        ""
      )}
    </svg>
  );
};
