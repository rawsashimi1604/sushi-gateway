import React from 'react';

interface TextAreaInputProps {
  id: string;
  name: string;
  placeholder?: string;
  value?: string;
  onChange?: (event: React.ChangeEvent<HTMLTextAreaElement>) => void;
  error?: string;
}

function TextAreaInput({ id, name, placeholder, value, onChange, error }: TextAreaInputProps) {
  return (
    <>
      <textarea
        id={id}
        name={name}
        className="font-sans tracking-wider indent-1 rounded-lg w-full bg-gray-100 border-sashimi-gray border-[0.5px] px-1 py-0.5 text-sm focus:outline-none h-[80px]"
        placeholder={placeholder}
        value={value}
        onChange={onChange}
      />
      {error && <span className="tracking-wider text-red-500 text-xs">{error}</span>}
    </>
  );
}

export default TextAreaInput;
