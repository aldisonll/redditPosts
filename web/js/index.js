const apiEndpoint = `http://${location.hostname}:3333/api`

const allRedditPosts = document.querySelector('[reddit-posts]')

const createRedditPostComponent = (postInfo) => {
    [link,
        author,
         title, 
         description, 
         url,
         numbersOfUpvotes, 
         numbersOfComments, 
         numberOfRewards
    ] = postInfo
    return `
<div post>

<span user>
<a post-link href="https://reddit.com${link}" target="_blank">Posted By u/${author}</a>
</span>

<span title>
${title}
</span>

<span description>
${description}
</span>

<span url>
    <a href="${url}" target="_blank">
        ${
            (url.length > 40) ? 
            url.substr(0, 40)  + '...' : url
        }
    </a>
</span>

<div info>
<span upvotes>${numbersOfUpvotes} ⬆️</span>
<span comments>${numbersOfComments} 📑</span>
<span rewards>${numberOfRewards} 💎</span>
</div>

<div actions>

<button like>Like</button>
<button comment>Comment</button>
<button rewards>Reward</button>

</div>

</div>
    `
}

const getPostsFromServer = async () => {
    const response = await fetch(apiEndpoint)
    const jsonData = await response.json()
    return jsonData
}

const getPostInfo = (postData) => {
    return new Array(
        postData.data.permalink,
        postData.data.author, 
        postData.data.title, 
        postData.data.description || "",
        postData.data.url_overridden_by_dest,
        postData.data.ups || 0, 
        postData.data.num_comments || 0, 
        postData.data.all_awardings.length ||  0
    )
}

async function main() {
    var allPosts = await getPostsFromServer()
    
    allPosts = allPosts.data.children

    allPosts.map((post) => {
        thisPostInfo = getPostInfo(post)
        allRedditPosts.innerHTML += createRedditPostComponent(thisPostInfo)
    })
}

main()