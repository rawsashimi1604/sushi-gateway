import React from "react";
import ReactDOM from "react-dom";

interface ModalProps {
  isOpen: boolean;
  onClose: () => void;
  title?: string;
  children: React.ReactNode | React.ReactNode[];
}

function Modal({ isOpen, onClose, title, children }: ModalProps) {
  if (!isOpen) return null;

  //   In React, ReactDOM.createPortal() is a method that allows you to render a component or element outside of its normal DOM hierarchy while keeping it part of React's rendering process. This is particularly useful for components like modals, tooltips, or dropdowns that need to appear outside their parent element's DOM structure but still need to maintain the React lifecycle.
  return ReactDOM.createPortal(
    <div
      className="fixed inset-0 bg-gray-800 bg-opacity-75 flex items-center justify-center z-50"
      onClick={onClose}
    >
      <div
        className="bg-white rounded-lg shadow-lg w-full max-w-lg p-6 relative"
        onClick={(e) => e.stopPropagation()} // Prevents closing the modal when clicking inside the modal
      >
        <div className="flex justify-between items-center border-b pb-3 mb-4">
          {title && (
            <h2 className="text-xl font-semibold text-gray-800">{title}</h2>
          )}
          <button
            onClick={onClose}
            className="text-gray-500 hover:text-gray-700 focus:outline-none focus:ring-2 focus:ring-gray-300 rounded-full"
          >
            X
          </button>
        </div>
        <div>{children}</div>
      </div>
    </div>,
    document.body
  );
}

export default Modal;
