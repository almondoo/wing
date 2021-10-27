import styled from 'styled-components';
import variable from '../../../utils/variable';

const Header = styled.header`
  position: fixed;
  width: 100%;
  height: ${variable.header.height};
  z-index: 10000;
`;

const TopBar = styled.div`
  width: 100%;
  height: 20px;
  background-color: ${variable.color.primaryDark};
`;

const BottomBar = styled.div`
  width: 100%;
  height: 60px;
  padding: 0 20px;
  background-color: ${variable.color.primary};
`;

const Title = styled.h1`
  font-size: 24px;
  font-weight: normal;
  line-height: 60px;
  color: ${variable.color.white};
`;

const exportDefault = {
  Header,
  TopBar,
  BottomBar,
  Title,
};

export default exportDefault;
