const MIN_FONT_SIZE = 12;
const MAX_FONT_SIZE = 24;
const FONT_SIZE_STEP = 2;
const FONT_SIZE_KEY = 'app-font-size';

export const getFontSize = (): number => {
  const savedSize = localStorage.getItem(FONT_SIZE_KEY);
  return savedSize ? parseInt(savedSize, 10) : 16;
};

export const setFontSize = (size: number): void => {
  const clampedSize = Math.min(Math.max(size, MIN_FONT_SIZE), MAX_FONT_SIZE);
  localStorage.setItem(FONT_SIZE_KEY, clampedSize.toString());
  document.documentElement.style.fontSize = `${clampedSize}px`;
};

export const increaseFontSize = (): void => {
  const currentSize = getFontSize();
  setFontSize(currentSize + FONT_SIZE_STEP);
};

export const decreaseFontSize = (): void => {
  const currentSize = getFontSize();
  setFontSize(currentSize - FONT_SIZE_STEP);
};

export const resetFontSize = (): void => {
  setFontSize(16);
}; 