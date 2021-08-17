const size = {
  mobile: '480px',
  tablet: '768px',
  pc: '1100px',
};

export const minDevice = {
  mobile: `(min-width: ${size.mobile})`,
  tablet: `(min-width: ${size.tablet})`,
  pc: `(min-width: ${size.pc})`,
};

export const maxDevice = {
  mobile: `(max-width: ${size.mobile})`,
  tablet: `(max-width: ${size.tablet})`,
  pc: `(max-width: ${size.pc})`,
};
