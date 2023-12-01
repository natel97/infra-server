import styled from "styled-components";
import { glowVariants, colorVariants } from "./generic";

type GlowProps = {
  glow?: keyof typeof glowVariants;
};

export const GlowingListItem = styled.div<GlowProps>`
  border-radius: 4px;
  background: ${colorVariants.primary};
  box-shadow: 0px 0px 2px 2px ${(props) => props.glow};
  color: white;
  padding: 8px;
  margin: 12px 0;
`;
