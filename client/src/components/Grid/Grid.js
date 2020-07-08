import React from 'react'
import Card from '../Card/Card.js'
import './Grid.css'
import { loadSubreddits } from "../../actions/index.js"
import { useDispatch, useSelector } from 'react-redux'



function Grid(props) {
    const {endpoint} = props
    const dispatch = useDispatch() 
    const name = useSelector(state => state.name)
    const subreddits = useSelector(state => state.subreddits)
    let cards = []
    if (subreddits !== []) {
        // eslint-disable-next-line array-callback-return
        cards = Object.entries(subreddits).map(([subName, keywords], index) => {
            return (
                //pass image as prop to card along with subreddits ect.
                <Card 
                    key={index} 
                    userName={name}
                    subName={subName} 
                    keywords={keywords} 
                    endpoint={endpoint}
                />
            )
        })
    } else {
        dispatch(loadSubreddits(name))
    }


    return (
        <div className="grid">
            {cards}
        </div>
    )
}

export default Grid