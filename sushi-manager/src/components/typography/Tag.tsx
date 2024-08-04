import React from 'react'

interface TagProps {
    value: string
}

function Tag({ value }: TagProps) {
    return (
        // TODO: change color scheme
        <span className='text-xs px-2.5 py-1 rounded-xl bg-neutral-500 text-white'>{value}</span>
    )
}

export default Tag