import React from "react";

export interface DashboardCardProps {
  children: React.ReactElement | React.ReactElement[] | React.ReactNode;
  className?: string;
}

function DashboardCard({ children, className }: DashboardCardProps) {
  return (
    <div className={`shadow-md bg-white rounded-lg ${className}`}>
      {children}
    </div>
  );
}

export default DashboardCard;
