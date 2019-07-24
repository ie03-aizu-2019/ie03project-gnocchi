import * as React from "react";

type PointProps = {
  x: number;
  y: number;
  size: number;
  color?: string;
  onMouseDown?: (e: React.MouseEvent) => void;
};

export default ({ color = "red", ...props }: PointProps) => {
  return (
    <circle
      onMouseDown={props.onMouseDown}
      cx={props.x}
      cy={props.y}
      r={props.size}
      fill={color}
    />
  );
};
