import React, { useState } from "react";
import Modal from "../../components/layout/Modal";

function PluginModal() {
  const [isPluginModalOpen, setIsPluginModalOpen] = useState(true);
  const openModal = () => setIsPluginModalOpen(true);
  const closeModal = () => setIsPluginModalOpen(false);

  return (
    <Modal isOpen={isPluginModalOpen} onClose={closeModal} title="Plugin Modal">
      <div>hello world</div>
    </Modal>
  );
}

export default PluginModal;
