import styled, { css } from 'styled-components';
import variable from '../../utils/variable';

const Wrapper = styled.div`
  width: 100%;
  height: calc((1em * 1.2) + 30px);
`;

const Field = styled.div`
  position: relative;
  width: 100%;
  height: 100%;
`;

const Label = styled.label<{ isInput: string | number; isFloat: boolean }>`
  position: absolute;
  z-index: 2;
  ${({ isInput, isFloat }) => {
    if (isInput || isFloat) {
      return css`
        top: 0;
        font-size: 12px;
        z-index: 4;
      `;
    } else {
      return css`
        top: 50%;
      `;
    }
  }}
  transform: translateY(-50%);
  left: 10px;
  max-width: 100%;
  padding: 0 3px;
  transition: all 0.3s;
  color: #333;
  background-color: #fff;
  border-radius: 5px;
  line-height: 1;
`;

const Input = styled.input`
  position: absolute;
  padding: 15px 10px;
  width: 100%;
  height: 100%;
  border: 1px solid grey;

  border-radius: 5px;
  &:hover {
    border: 1px solid #000;
  }
  &:focus {
    border: 2px solid ${variable.color.primary};

    ~ ${Label} {
      font-size: 12px;
      top: 0;
      z-index: 4;
      color: ${variable.color.primary};
    }
  }
`;

const EyeIcon = styled.span`
  position: absolute;
  top: 50%;
  right: 15px;
  transform: translateY(-50%);
  display: block;
  width: 20px;
  height: 20px;
  background-image: url('/images/svg/eye.svg');
  background-repeat: no-repeat;
  z-index: 4;
  cursor: pointer;
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
  EyeIcon,
  Required,
};

export default exportDefault;
