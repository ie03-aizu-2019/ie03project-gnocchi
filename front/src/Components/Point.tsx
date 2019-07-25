import * as React from "react";

import { Place } from "../Reducer";

type PointProps = {
  place: Place;
  size: number;
  color?: string;
  index?: number;
  onMouseDown?: (e: React.MouseEvent) => void;
};

export default ({ color = "red", place, ...props }: PointProps) => {
  return (
    <>
      <circle
        onMouseDown={props.onMouseDown}
        cx={place.x}
        cy={place.y}
        r={props.size}
        fill={color}
      />
      <text
        x={place.x - 2 * props.size}
        y={place.y - 2 * props.size}
        fontSize={props.size * 2}
      >
        {props.index}
      </text>
    </>
  );
};
