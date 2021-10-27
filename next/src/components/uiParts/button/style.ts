import styled, { css } from 'styled-components';

type ButtonType = {
  fullWidth: boolean;
  variant: string;
  color: string;
};

const Button = styled.button<ButtonType>`
  display: block;
  padding: 8px 25px;
  text-align: center;
  transition: all 0.2s;
  font-size: 1.2rem;
  color: #fff;

  ${({ fullWidth }) => {
    if (fullWidth) {
      return css`
        width: 100%;
      `;
    }
  }}

  ${({ variant, color }) => {
    if (variant == 'contained') {
      return css`
        background-color: ${color};
        border-radius: 10px;
        box-shadow: 0px 3px 3px 1px rgba(0, 0, 0, 0.3);
      `;
    } else {
      return css`
        &:hover {
          text-decoration-line: underline;
        }
      `;
    }
  }}

  &:hover {
    opacity: 0.9;
  }

  &:active {
    box-shadow: none;
  }
`;

const exportDefault = {
  Button,
};

export default exportDefault;
