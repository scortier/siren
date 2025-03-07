package slacknotifier

import (
	"github.com/odpf/siren/domain"
	"github.com/pkg/errors"
	"testing"

	"github.com/odpf/siren/mocks"
	"github.com/slack-go/slack"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type SlackHTTPClientTestSuite struct {
	suite.Suite
}

func TestHTTP(t *testing.T) {
	suite.Run(t, new(SlackHTTPClientTestSuite))
}

func (s *SlackHTTPClientTestSuite) SetupTest() {}

func (s *SlackHTTPClientTestSuite) TestSlackHTTPClient_Notify() {
	oldServiceCreator := newService
	mockedSlackService := &mocks.SlackService{}
	newService = func(string) domain.SlackService {
		return mockedSlackService
	}
	defer func() { newService = oldServiceCreator }()
	testNotifierClient := SlackNotifierClient{Slacker: mockedSlackService}

	s.Run("should notify user identified by their email", func() {
		mockedSlackService.On("GetUserByEmail", "foo@odpf.io").
			Return(&slack.User{ID: "U20"}, nil).Once()
		mockedSlackService.On("SendMessage", "U20",
			mock.AnythingOfType("slack.MsgOption"), mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil).Once()
		dummyMessage := &SlackMessage{
			ReceiverName: "foo@odpf.io",
			ReceiverType: "user",
			Message:      "random text",
		}
		err := testNotifierClient.Notify(dummyMessage, "foo_bar")
		s.Nil(err)
	})

	s.Run("should return error if notifying user fails", func() {
		mockedSlackService.On("GetUserByEmail", "foo@odpf.io").
			Return(&slack.User{ID: "U20"}, nil).Once()
		mockedSlackService.On("SendMessage", "U20",
			mock.AnythingOfType("slack.MsgOption"),
			mock.AnythingOfType("slack.MsgOption"),
		).Return("", "", "", errors.New("random error")).Once()

		dummyMessage := &SlackMessage{
			ReceiverName: "foo@odpf.io",
			ReceiverType: "user",
			Message:      "random text",
		}
		err := testNotifierClient.Notify(dummyMessage, "foo_bar")
		s.EqualError(err, "failed to send message to foo@odpf.io: random error")
	})

	s.Run("should return error if user lookup by email fails", func() {
		mockedSlackService.On("GetUserByEmail", "foo@odpf.io").
			Return(nil, errors.New("users_not_found")).Once()

		dummyMessage := &SlackMessage{
			ReceiverName: "foo@odpf.io",
			ReceiverType: "user",
			Message:      "random text",
		}
		err := testNotifierClient.Notify(dummyMessage, "foo_bar")
		s.EqualError(err, "failed to get id for foo@odpf.io: users_not_found")
	})

	s.Run("should return error if user lookup by email fails", func() {
		mockedSlackService.On("GetUserByEmail", "foo@odpf.io").
			Return(nil, errors.New("random error")).Once()

		dummyMessage := &SlackMessage{
			ReceiverName: "foo@odpf.io",
			ReceiverType: "user",
			Message:      "random text",
		}
		err := testNotifierClient.Notify(dummyMessage, "foo_bar")
		s.EqualError(err, "random error")
	})

	s.Run("should notify if part of the channel", func() {
		mockedSlackService.On("GetJoinedChannelsList").Return([]slack.Channel{
			{GroupConversation: slack.GroupConversation{
				Name:         "foo",
				Conversation: slack.Conversation{ID: "C01"}},
			}, {GroupConversation: slack.GroupConversation{
				Name:         "bar",
				Conversation: slack.Conversation{ID: "C02"}},
			}}, nil).Once()

		mockedSlackService.On("SendMessage", "C01",
			mock.AnythingOfType("slack.MsgOption"),
			mock.AnythingOfType("slack.MsgOption"),
		).Return("", "", "", nil).Once()

		dummyMessage := &SlackMessage{
			ReceiverName: "foo",
			ReceiverType: "channel",
			Message:      "random text",
		}
		err := testNotifierClient.Notify(dummyMessage, "foo_bar")
		s.Nil(err)
		mockedSlackService.AssertExpectations(s.T())
	})

	s.Run("should return error if not part of the channel", func() {
		mockedSlackService.On("GetJoinedChannelsList").Return([]slack.Channel{
			{GroupConversation: slack.GroupConversation{
				Name:         "foo",
				Conversation: slack.Conversation{ID: "C01"}},
			}, {GroupConversation: slack.GroupConversation{
				Name:         "bar",
				Conversation: slack.Conversation{ID: "C02"}},
			}}, nil).Once()

		mockedSlackService.On("SendMessage", "C01",
			mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil).Once()

		dummyMessage := &SlackMessage{
			ReceiverName: "baz",
			ReceiverType: "channel",
			Message:      "random text",
		}
		err := testNotifierClient.Notify(dummyMessage, "foo_bar")
		s.EqualError(err, "app is not part of the channel baz")
	})

	s.Run("should return error failed to fetch joined channels list", func() {
		mockedSlackService.On("GetJoinedChannelsList").
			Return(nil, errors.New("random error")).Once()
		mockedSlackService.On("SendMessage", "C01",
			mock.AnythingOfType("slack.MsgOption")).Return("", "", "", nil).Once()

		dummyMessage := &SlackMessage{
			ReceiverName: "baz",
			ReceiverType: "channel",
			Message:      "random text",
		}
		err := testNotifierClient.Notify(dummyMessage, "foo_bar")
		s.EqualError(err, "failed to fetch joined channel list: random error")
	})
}
