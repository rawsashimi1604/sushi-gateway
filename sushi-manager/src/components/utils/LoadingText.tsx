import React from 'react';

import LoadingSpinner from './LoadingSpinner';

interface LoadingTextProps {
  text: string;
  spinnerSize?: number;
}

function LoadingText({ text, spinnerSize = 14 }: LoadingTextProps) {
  return (
    <div className="flex flex-row items-center gap-2 h-96 justify-center">
      <span className="font-lora tracking-wider text-gray-600 text-sm">{text}</span>
      <LoadingSpinner size={spinnerSize} />
    </div>
  );
}

export default LoadingText;
