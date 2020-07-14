import React from 'react'
import './Tag.css'
function Tag(props) {
    const { word } = props
    return (
        <div className="tag">{word}</div>
    )
}

export default Tag