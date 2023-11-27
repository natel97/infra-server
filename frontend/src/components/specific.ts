import styled from "styled-components";
import { colorVariants } from "./generic";

type GlowProps = {
  glow?: keyof typeof colorVariants;
};

export const GlowingListItem = styled.div<GlowProps>`
  border-radius: 4px;
  background: ${colorVariants.primary};
  box-shadow: 0px 0px 2px 2px ${(props) => props.glow};
  color: white;
  padding: 8px;
  margin: 12px 0;
`;
