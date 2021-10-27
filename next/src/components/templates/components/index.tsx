import { ChangeEvent, useState } from 'react';
import Style from './style';
import TextField from '../../uiParts/textField/index';
import TextArea from '../../uiParts/textarea/index';
import Button from '../../uiParts/button/index';
import Icon from '../../uiParts/icon/index';
import Image from '../../uiParts/image/index';
import Card from '../../uiParts/card/index';
import Link from '../../uiParts/link/index';
import Modal from '../../originals/modal/index';
import ModalContext from '../../originals/modal/context';

const Components = (): JSX.Element => {
  const [textFieldValue, setTextFieldValue] = useState<string>('');
  const [isModalOpen, setIsModalOpen] = useState<boolean>(false);

  return (
    <Style.Wrapper>
      <Style.Title>コンポーネント一覧</Style.Title>
      <Style.ComponentTypeTitle>UI Parts</Style.ComponentTypeTitle>
      <Style.ComponentTypeWrap>
        <Style.ComponentWrap>
          <Style.ComponentTitle>TextField</Style.ComponentTitle>
          <TextField
            value={textFieldValue}
            onChange={(e: ChangeEvent<HTMLInputElement>) => setTextFieldValue(e.target.value)}
          >
            テキストフィールド
          </TextField>
        </Style.ComponentWrap>
        <Style.ComponentWrap>
          <Style.ComponentTitle>TextArea</Style.ComponentTitle>
          <TextArea
            value={textFieldValue}
            onChange={(e: ChangeEvent<HTMLTextAreaElement>) => setTextFieldValue(e.target.value)}
          >
            テキストフィールド
          </TextArea>
        </Style.ComponentWrap>
        <Style.ComponentWrap>
          <Style.ComponentTitle>Button</Style.ComponentTitle>
          <Button onClick={() => console.log('button')}>ボタン</Button>
        </Style.ComponentWrap>
        <Style.ComponentWrap>
          <Style.ComponentTitle>Icon</Style.ComponentTitle>
          <Icon src="https://picsum.photos/200/200" alt="icon" size={100} />
        </Style.ComponentWrap>
        <Style.ComponentWrap>
          <Style.ComponentTitle>Image</Style.ComponentTitle>
          <Image src="https://picsum.photos/300/500" alt="icon" width="300px" height="300px" />
        </Style.ComponentWrap>
        <Style.ComponentWrap>
          <Style.ComponentTitle>Card</Style.ComponentTitle>
          <Card>カードこの中にいろいろ要素を追加できるよ！</Card>
        </Style.ComponentWrap>
        <Style.ComponentWrap>
          <Style.ComponentTitle>Link</Style.ComponentTitle>
          <Link href="/components">リンク先設定</Link>
        </Style.ComponentWrap>
      </Style.ComponentTypeWrap>

      <Style.ComponentTypeTitle>Originals</Style.ComponentTypeTitle>
      <Style.ComponentTypeWrap>
        <Style.ComponentWrap>
          <Style.ComponentTitle>Modal</Style.ComponentTitle>
          <ModalContext.Provider value={{ isOpen: isModalOpen, setIsOpen: setIsModalOpen }}>
            <Modal>モーダルの中にいろいろな要素を追加できるよ！</Modal>
          </ModalContext.Provider>
          <Button onClick={() => setIsModalOpen(true)}>Open</Button>
        </Style.ComponentWrap>
      </Style.ComponentTypeWrap>
    </Style.Wrapper>
  );
};

export default Components;
