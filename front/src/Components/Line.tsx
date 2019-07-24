import * as React from "react";

type LineProps = {
  from: { x: number; y: number };
  to: { x: number; y: number };
  width: number;
  color?: string;
};

export default ({ color = "black", ...props }: LineProps) => {
  return (
    <line
      x1={props.from.x}
      y1={props.from.y}
      x2={props.to.x}
      y2={props.to.y}
      stroke={color}
      strokeWidth={props.width}
    />
  );
};
