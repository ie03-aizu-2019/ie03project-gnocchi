import * as React from "react";
import styled from "styled-components";

export default styled.div<{ direction: string }>`
  display: flex;
  flex-direction: ${props => props.direction};
`;
