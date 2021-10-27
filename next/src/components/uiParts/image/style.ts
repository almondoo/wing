import styled from 'styled-components';

const Wrapper = styled.div<{ setWidth: string; setHeight: string }>`
  position: relative;
  width: ${({ setWidth }) => setWidth};
  height: ${({ setHeight }) => setHeight};
`;

const exportDefault = {
  Wrapper,
};

export default exportDefault;
