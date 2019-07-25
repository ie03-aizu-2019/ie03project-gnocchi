import * as React from "react";
import styled from "styled-components";

type GridProps = {
  rows: string[];
  columns: string[];
  gap?: string;
};

export default styled.div<GridProps>`
  display: grid;
  grid-template-rows: ${props => props.rows.join(" ")};
  grid-template-columns: ${props => props.columns.join(" ")};
  gap: ${props => props.gap || "0px"};
`;
