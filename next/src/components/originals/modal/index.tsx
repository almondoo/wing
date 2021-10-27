import { useContext, ReactNode } from 'react';
import Style from './style';
import Context from './context';
import Button from '../../uiParts/button/index';

type Props = {
  children: ReactNode;
};

const Modal = ({ children, ...props }: Props): JSX.Element => {
  const { isOpen, setIsOpen } = useContext(Context);

  return (
    <Style.Wrapper open={isOpen} {...props}>
      <Style.Modal>
        {children}
        <Style.ButtonGroup justifyContent="flex-end">
          <Button color="grey" onClick={() => setIsOpen(false)}>
            Close
          </Button>
        </Style.ButtonGroup>
      </Style.Modal>
    </Style.Wrapper>
  );
};

export default Modal;
