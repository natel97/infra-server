import styled from "styled-components";

export const colorVariants = {
  red: "#C53102",
  green: "#0E6100",
  blue: "#0A2A66",
  yellow: "#BFDE00",
  darkBlue: "#000714",
  primary: "#18171B",
  white: "white",
};

const fontSize = {
  sm: "12px",
  md: "18px",
  lg: "24px",
  xl: "32px",
};

const edgeSize = {
  xs: "2px",
  sm: "4px",
  md: "6px",
  lg: "10px",
  xl: "14px",
};

type Props = {
  color?: keyof typeof colorVariants;
  background?: keyof typeof colorVariants;
  fontSize?: keyof typeof fontSize;
  margin?: keyof typeof edgeSize;
  padding?: keyof typeof edgeSize;
};

export const Button = styled.button<Props>`
  border-radius: 4px;
  background: ${(props) => colorVariants[props.background || "primary"]};
  color: ${(props) => colorVariants[props.color || "white"]};
  font-size: ${(props) => fontSize[props.fontSize || "md"]};
  ${(props) => props.margin && `margin: ${edgeSize[props.margin]}`};
  ${(props) => props.padding && `padding: ${edgeSize[props.padding]}`};
`;
