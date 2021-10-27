import styled, { css } from 'styled-components';
import variable from '../../utils/variable';

const Wrapper = styled.div`
  width: 100%;
  background-color: #fff;
`;

const Field = styled.div<{ height: string }>`
  position: relative;
  width: 100%;
  height: ${({ height }) => height};
`;

const Input = styled.textarea<{ height: string }>`
  position: absolute;
  resize: none;
  padding: 15px 10px;
  width: 100%;
  height: ${({ height }) => height};
  border: 1px solid grey;
  border-radius: 5px;
  &:hover {
    border: 1px solid #000;
  }
  &:focus {
    border: 2px solid ${variable.color.primary};
    ~ label {
      font-size: 12px;
      top: 0;
      z-index: 4;
      color: ${variable.color.primary};
    }
  }
`;

const Label = styled.label<{ isInput: string }>`
  position: absolute;
  z-index: 2;
  ${({ isInput }) => {
    if (isInput) {
      return css`
        top: 0;
        font-size: 12px;
        z-index: 4;
      `;
    } else {
      return css`
        top: 20px;
      `;
    }
  }}
  transform: translateY(-50%);
  left: 10px;
  max-width: 100%;
  padding: 0 3px;
  transition: all 0.3s;
  background-color: #fff;
  border-radius: 5px;
`;

const Required = styled.span`
  margin-left: 5px;
  color: red;
`;

const exportDefault = {
  Wrapper,
  Field,
  Input,
  Label,
  Required,
};

export default exportDefault;
