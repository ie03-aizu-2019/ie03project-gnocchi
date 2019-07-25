import * as React from "react";

type LineProps = {
  from: { x: number; y: number };
  to: { x: number; y: number };
  width: number;
  color?: string;
};

export default ({ color = "black", ...props }: LineProps) => {
  const xLen = Math.abs(props.from.x - props.to.x);
  const yLen = Math.abs(props.from.y - props.to.y);

  const len = Math.round(Math.sqrt(xLen * xLen + yLen * yLen) * 100) / 100;

  return (
    <>
      <line
        x1={props.from.x}
        y1={props.from.y}
        x2={props.to.x}
        y2={props.to.y}
        stroke={color}
        strokeWidth={props.width}
      />
      <text
        x={Math.min(props.from.x, props.to.x) + xLen / 2}
        y={Math.min(props.from.y, props.to.y) + xLen / 2}
        fontSize={props.width * 5}
      >
        {len}
      </text>
    </>
  );
};
