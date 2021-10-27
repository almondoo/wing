import { createContext, SetStateAction } from 'react';
import type { Dispatch } from 'react';

type Type = {
  isOpen: boolean;
  setIsOpen: Dispatch<SetStateAction<boolean>>;
};

export default createContext<Type>({ isOpen: false, setIsOpen: () => {} });
