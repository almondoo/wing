import styled from 'styled-components';
import variable from '../../../utils/variable';

const Footer = styled.footer`
  width: 100%;
  background-color: ${variable.color.primary};
  color: ${variable.color.white};
`;

const Inner = styled.div`
  max-width: 1020px;
  margin: 0 auto;
  padding: 30px 20px;
`;

const exportDefault = {
  Footer,
  Inner,
};

export default exportDefault;
