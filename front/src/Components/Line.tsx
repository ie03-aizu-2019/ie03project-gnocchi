import * as React from "react";

type LineProps = {
  from: { x: number; y: number };
  to: { x: number; y: number };
  width: number;
  color?: string;
  isShowLength: boolean;
};

export default ({ color = "black", ...props }: LineProps) => {
  const xLen = Math.abs(props.from.x - props.to.x);
  const yLen = Math.abs(props.from.y - props.to.y);
  const len = Math.round(Math.sqrt(xLen * xLen + yLen * yLen) * 100) / 100;

  const rand = Math.random();

  return (
    <>
      <path
        d={`M${props.from.x},${props.from.y} L${props.to.x},${props.to.y}`}
        stroke={color}
        strokeWidth={props.width}
      />
      {props.isShowLength ? (
        <>
          <defs>
            <path
              d={`M${props.from.x},${props.from.y - 2 * props.width} L${
                props.to.x
              },${props.to.y - 2 * props.width}`}
              id={`line${rand}`}
            />
          </defs>
          <text fontSize={props.width * 5}>
            <textPath
              href={`#line${rand}`}
              textAnchor="middle"
              startOffset="50%"
            >
              {len}
            </textPath>
          </text>
        </>
      ) : (
        ""
      )}
    </>
  );
};
