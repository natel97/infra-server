import styled, { createGlobalStyle } from "styled-components";

export const colorVariants = {
  red: "#C53102",
  green: "#0E6100",
  blue: "#0A2A66",
  yellow: "#BFDE00",
  darkBlue: "#000714",
  primary: "#18171B",
  white: "white",
  gray: "#606060",
};

export const glowVariants = {
  red: "#C5310233",
  green: "#0E610033",
  blue: "#0A2A6633",
  yellow: "#BFDE0033",
  darkBlue: "#00071433",
  primary: "#18171B33",
  white: "#FFFFFF33",
  gray: "#60606033",
};

export const GlobalStyles = createGlobalStyle`
  html {
    background: ${colorVariants.darkBlue};
    height: 100%;
    width: 100%;
  }

  body {
    height: 100%;
    width: 100%;
    margin: 0;
  }

  h1, h2, h3, h4, h5 {
    margin: 0;
    color: ${colorVariants.white}
  }
  `;

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
