import styled from 'styled-components';

const ButtonWrap = styled.div<{ justifyContent: string }>`
  display: flex;
  justify-content: ${({ justifyContent }) => justifyContent};
  align-items: center;
  width: 100%;
  button {
    margin-left: 20px;
    &:first-of-type {
      margin-left: 0;
    }
  }
`;

const exportDefault = {
  ButtonWrap,
};

export default exportDefault;
