import styled, { css } from 'styled-components';

const Link = styled.a<{ isUnderline: boolean }>`
  ${({ isUnderline }) => {
    if (!isUnderline) {
      return css`
        &:hover {
          text-decoration: none;
        }
      `;
    }
  }}
`;

const exportDefault = {
  Link,
};

export default exportDefault;
