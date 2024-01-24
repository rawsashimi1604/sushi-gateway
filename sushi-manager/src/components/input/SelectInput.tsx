import React from 'react';

interface SelectInputProps {
  options: string[];
  onChange?: (value: string) => void;
  error?: string;
  value?: string;
}

function SelectInput({ options, onChange, error, value }: SelectInputProps) {
  return (
    <div>
      <select
        className="font-sans tracking-wider indent-1 rounded-lg w-full bg-gray-100 border-sashimi-gray border-[0.5px] px-1 py-0.5 text-sm focus:outline-none"
        onChange={(e) => onChange && onChange(e.target.value)}
        value={value}
      >
        <option value="">Select an option</option>
        {options?.map((option) => {
          return (
            <option value={option} key={option}>
              {option}
            </option>
          );
        })}
      </select>
      {error && <span className="tracking-wider text-red-500 text-xs">{error}</span>}
    </div>
  );
}

export default SelectInput;
