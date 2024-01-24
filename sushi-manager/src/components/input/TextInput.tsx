import React from 'react';

interface TextInputProps {
  id: string;
  name: string;
  placeholder?: string;
  value?: string;
  onChange?: (event: React.ChangeEvent<HTMLInputElement>) => void;
  error?: string;
  type?: 'text' | 'password';
}

function TextInput({ id, name, placeholder, value, onChange, error, type = 'text' }: TextInputProps) {
  return (
    <>
      <input
        id={id}
        name={name}
        type={type}
        className="font-sans tracking-wider indent-1 rounded-lg w-full bg-gray-100 border-sashimi-gray border-[0.5px] px-1 py-0.5 text-sm focus:outline-none"
        placeholder={placeholder}
        value={value}
        onChange={onChange}
      />
      {error && <span className="tracking-wider text-red-500 text-xs">{error}</span>}
    </>
  );
}

export default TextInput;
