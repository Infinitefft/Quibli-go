import React, { useState, useEffect, useCallback } from 'react';
import { throttle } from '@/utils/Throttle';

interface BackToTopProps {
  targetRef: React.RefObject<HTMLDivElement>;
  threshold?: number;
  right?: number;
  bottom?: number;
  tabBarHeight?: number;
}

const BackToTop: React.FC<BackToTopProps> = ({ 
  targetRef,
  threshold = 400, 
  right = 16, 
  bottom = 20,
  tabBarHeight = 0 
}) => {
  const [visible, setVisible] = useState(false);

  const checkScroll = useCallback(
    throttle(() => {
      if (targetRef.current) {
        setVisible(targetRef.current.scrollTop > threshold);
      }
    }, 200),
    [threshold, targetRef]
  );

  useEffect(() => {
    const el = targetRef.current;
    if (el) {
      el.addEventListener('scroll', checkScroll);
      return () => el.removeEventListener('scroll', checkScroll);
    }
  }, [checkScroll, targetRef]);

  const handleToTop = () => {
    targetRef.current?.scrollTo({
      top: 0,
      behavior: 'smooth'
    });
  };

  if (!visible) return null;

  return (
    <div
      onClick={handleToTop}
      style={{
        position: 'fixed',
        right: `${right}px`,
        bottom: `${tabBarHeight + bottom}px`,
        width: '42px',
        height: '42px',
        borderRadius: '50%',
        backgroundColor: '#ffffff',
        boxShadow: '0 3px 10px rgba(0,0,0,0.12)',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
        zIndex: 999,
        cursor: 'pointer',
        WebkitTapHighlightColor: 'transparent'
      }}
    >
      <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="#666" strokeWidth="2.5" strokeLinecap="round" strokeLinejoin="round">
        <path d="M18 15l-6-6-6 6"/>
      </svg>
    </div>
  );
};

export default BackToTop;