import Style from './style';

const Header = (): JSX.Element => {
  return (
    <Style.Header>
      <Style.TopBar />
      <Style.BottomBar>
        <Style.Title>タイトル</Style.Title>
      </Style.BottomBar>
    </Style.Header>
  );
};

export default Header;
