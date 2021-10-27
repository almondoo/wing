import styled from 'styled-components';
import MaterialCard from '../../uiParts/card/index';
import MaterialButtonGroup from '../../uiParts/buttonGroup/index';

const Wrapper = styled.div<{ open: boolean }>`
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  z-index: 10000;
  background-color: rgba(0, 0, 0, 0.4);
  visibility: ${({ open }) => (open ? 'visible' : 'hidden')};
  transition: all 0.3s;
  padding: 0 15px;
`;

const Modal = styled(MaterialCard)`
  position: relative;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  max-width: 1000px;
  max-height: 800px;
  padding: 30px;
  overflow-y: scroll;
`;

const ButtonGroup = styled(MaterialButtonGroup)`
  margin-top: 20px;
`;

const exportDefault = {
  Wrapper,
  Modal,
  ButtonGroup,
};

export default exportDefault;
