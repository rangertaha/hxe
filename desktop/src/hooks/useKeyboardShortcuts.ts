import { useEffect } from 'react';
import { increaseFontSize, decreaseFontSize, resetFontSize } from '../utils/fontSize';

export const useKeyboardShortcuts = () => {
  useEffect(() => {
    const handleKeyDown = (event: KeyboardEvent) => {
      // Check if Ctrl/Cmd key is pressed
      if (event.ctrlKey || event.metaKey) {
        switch (event.key) {
          case '+':
            event.preventDefault();
            increaseFontSize();
            break;
          case '-':
            event.preventDefault();
            decreaseFontSize();
            break;
          case '0':
            event.preventDefault();
            resetFontSize();
            break;
        }
      }
    };

    window.addEventListener('keydown', handleKeyDown);
    return () => window.removeEventListener('keydown', handleKeyDown);
  }, []);
}; 