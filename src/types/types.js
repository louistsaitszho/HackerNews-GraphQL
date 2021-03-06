const {
  GraphQLObjectType,
  GraphQLList,
  GraphQLID,
  GraphQLInt,
} = require('graphql');
const {
  resolveTopStories,
  resolveBestStories,
  resolveNewStories,
  resolveStoryByID,
  resolveJobs,
} = require('../resolvers/resolvers.js');
const userResolvers = require('../resolvers/userResolvers.js');
const storyResolvers = require('../resolvers/storyResolvers.js');
const basicTypes = require('./basicTypes.js');

const QueryType = new GraphQLObjectType({
  name: 'Query',
  fields: () => ({
    topStories: {
      type: new GraphQLList(basicTypes.StoryListItemType),
      args: {
        limit: { type: GraphQLInt },
        offset: { type: GraphQLInt },
      },
      resolve: (_, { limit, offset }) => resolveTopStories(limit, offset),
    },
    bestStories: {
      type: new GraphQLList(basicTypes.MaybeStoryType),
      args: {
        limit: { type: GraphQLInt },
        offset: { type: GraphQLInt },
      },
      resolve: (_, { limit, offset }) => resolveBestStories(limit, offset),
    },
    newStories: {
      type: new GraphQLList(basicTypes.MaybeStoryType),
      args: {
        limit: { type: GraphQLInt },
        offset: { type: GraphQLInt },
      },
      resolve: (_, { limit, offset }) => resolveNewStories(limit, offset),
    },
    story: {
      type: basicTypes.MaybeStoryType,
      args: {
        id: { type: GraphQLID },
      },
      resolve: (_, { id }) => storyResolvers.storyByID(id),
    },
    user: {
      type: basicTypes.UserType,
      args: {
        id: { type: GraphQLID },
      },
      resolve: (_, { id }) => userResolvers.userByID(id),
    },
    jobs: {
      type: new GraphQLList(basicTypes.JobType),
      args: {
        limit: { type: GraphQLInt },
        offset: { type: GraphQLInt },
      },
      resolve: (_, { limit, offset }) => resolveJobs(limit, offset),
    },
  }),
});

module.exports = {
  QueryType,
};
